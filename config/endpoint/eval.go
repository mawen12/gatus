package endpoint

import (
	"strconv"
	"strings"

	"github.com/TwiN/gatus/v5/config/gontext"
	"github.com/TwiN/gatus/v5/jsonpath"
)

// EvalPlaceholder parses and evaluates a placeholder expression.
//
// Examples:
//   - [STATUS]
//   - len([BODY].items)
//   - has([CONTEXT].user_id)
func EvalPlaceholder(input string, result *Result, ctx *gontext.Gontext) (string, error) {
	expr := parseEvalExpression(input)
	return expr.Eval(result, ctx)
}

type evalExpression interface {
	Eval(result *Result, ctx *gontext.Gontext) (string, error)
}

type expressionKind int

const (
	expressionLiteral expressionKind = iota
	expressionStatus
	expressionIP
	expressionResponseTime
	expressionDNSRCode
	expressionConnected
	expressionCertificateExpiration
	expressionDomainExpiration
	expressionBody
	expressionBodyPath
	expressionContextPath
	expressionUnknown
)

type literalExpression struct {
	value string
}

func (e literalExpression) Eval(_ *Result, _ *gontext.Gontext) (string, error) {
	return e.value, nil
}

type placeholderExpression struct {
	original    string
	placeholder string
	fn          functionType
	kind        expressionKind
}

func (e placeholderExpression) Eval(result *Result, ctx *gontext.Gontext) (string, error) {
	switch e.kind {
	case expressionStatus:
		return formatWithFunction(strconv.Itoa(result.HTTPStatus), e.fn), nil
	case expressionIP:
		return formatWithFunction(result.IP, e.fn), nil
	case expressionResponseTime:
		return formatWithFunction(strconv.FormatInt(result.Duration.Milliseconds(), 10), e.fn), nil
	case expressionDNSRCode:
		return formatWithFunction(result.DNSRCode, e.fn), nil
	case expressionConnected:
		return formatWithFunction(strconv.FormatBool(result.Connected), e.fn), nil
	case expressionCertificateExpiration:
		return formatWithFunction(strconv.FormatInt(result.CertificateExpiration.Milliseconds(), 10), e.fn), nil
	case expressionDomainExpiration:
		return formatWithFunction(strconv.FormatInt(result.DomainExpiration.Milliseconds(), 10), e.fn), nil
	case expressionBody:
		return evalBody(result, e.fn)
	case expressionBodyPath:
		return resolveJSONPathPlaceholder(e.placeholder, e.fn, e.original, result)
	case expressionContextPath:
		if ctx == nil {
			return evalUnknownWithFunction(e)
		}
		return resolveContextPlaceholder(e.placeholder, e.fn, e.original, ctx)
	default:
		if e.fn != noFunction {
			return evalUnknownWithFunction(e)
		}
		return e.original, nil
	}
}

func evalUnknownWithFunction(e placeholderExpression) (string, error) {
	if e.fn == functionHas {
		return "false", nil
	}
	return e.original + " " + InvalidConditionElementSuffix, nil
}

func evalBody(result *Result, fn functionType) (string, error) {
	body := strings.TrimSpace(string(result.Body))
	if fn == functionHas {
		return strconv.FormatBool(len(body) > 0), nil
	}
	if fn == functionLen {
		// If body is valid JSON, return collection/object length, else string length.
		_, resolvedLength, err := jsonpath.Eval("", result.Body)
		if err == nil {
			return strconv.Itoa(resolvedLength), nil
		}
		return strconv.Itoa(len(body)), nil
	}
	return body, nil
}

func parseEvalExpression(input string) evalExpression {
	trimmed := strings.TrimSpace(input)
	fn, innerPlaceholder := extractFunctionWrapper(trimmed)
	placeholder := innerPlaceholder
	uppercasePlaceholder := strings.ToUpper(placeholder)

	kind := classifyExpressionKind(uppercasePlaceholder)
	if kind == expressionUnknown && fn == noFunction {
		return literalExpression{value: trimmed}
	}

	return placeholderExpression{
		original:    trimmed,
		placeholder: placeholder,
		fn:          fn,
		kind:        kind,
	}
}

func classifyExpressionKind(uppercasePlaceholder string) expressionKind {
	switch uppercasePlaceholder {
	case StatusPlaceholder:
		return expressionStatus
	case IPPlaceholder:
		return expressionIP
	case ResponseTimePlaceholder:
		return expressionResponseTime
	case DNSRCodePlaceholder:
		return expressionDNSRCode
	case ConnectedPlaceholder:
		return expressionConnected
	case CertificateExpirationPlaceholder:
		return expressionCertificateExpiration
	case DomainExpirationPlaceholder:
		return expressionDomainExpiration
	case BodyPlaceholder:
		return expressionBody
	}

	if strings.HasPrefix(uppercasePlaceholder, BodyPlaceholder+".") || strings.HasPrefix(uppercasePlaceholder, BodyPlaceholder+"[") {
		return expressionBodyPath
	}
	if strings.HasPrefix(uppercasePlaceholder, ContextPlaceholder) {
		return expressionContextPath
	}

	return expressionUnknown
}
