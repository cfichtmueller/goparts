// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package e

const (
	AttrHXGet       = "hx-get"
	AttrHXPost      = "hx-post"
	AttrHXPushUrl   = "hx-push-url"
	AttrHXSelect    = "hx-select"
	AttrHXSelectOob = "hx-select-oob"
	AttrHXSwap      = "hx-swap"
	AttrHXSwapOob   = "hx-swap-oob"
	AttrHXTarget    = "hx-target"
	AttrHXTrigger   = "hx-trigger"
	AttrHXVals      = "hx-vals"

	AttrHXBoost       = "hx-boost"
	AttrHXConfirm     = "hx-confirm"
	AttrHXDelete      = "hx-delete"
	AttrHXDisable     = "hx-disable"
	AttrHXDisabledElt = "hx-disabled-elt"
	AttrHXDisinherit  = "hx-disinherit"
	AttrHXEncoding    = "hx-encoding"
	AttrHXExt         = "hx-ext"
	AttrHXHeaders     = "hx-headers"
	AttrHXHistory     = "hx-history"
	AttrHXHistoryElt  = "hx-history-elt"
	AttrHXInclude     = "hx-include"
	AttrHXIndicator   = "hx-indicator"
	AttrHXInherit     = "hx-inherit"
	AttrHXParams      = "hx-params"
	AttrHXPatch       = "hx-patch"
	AttrHXPreserve    = "hx-preserve"
	AttrHXPrompt      = "hx-prompt"
	AttrHXPut         = "hx-put"
	AttrHXReplaceUrl  = "hx-replace-url"
	AttrHXRequest     = "hx-request"
	AttrHXSync        = "hx-sync"
	AttrHXValidate    = "hx-validate"

	HXEvtAbort                 = "htmx:abort"
	HXEvtAfterOnLoad           = "htmx:afterOnLoad"
	HXEvtAfterProcessNode      = "htmx:afterProcessNode"
	HXEvtAfterRequest          = "htmx:afterRequest"
	HXEvtAfterSettle           = "htmx:afterSettle"
	HXEvtAfterSwap             = "htmx:afterSwap"
	HXEvtBeforeCleanupElement  = "htmx:beforeCleanupElement"
	HXEvtBeforeOnLoad          = "htmx:beforeOnLoad"
	HXEvtBeforeProcessNode     = "htmx:beforeProcessNode"
	HXEvtBeforeRequest         = "htmx:beforeRequest"
	HXEvtBeforeSwap            = "htmx:beforeSwap"
	HXEvtBeforeSend            = "htmx:beforeSend"
	HXEvtBeforeTransition      = "htmx:beforeTransition"
	HXEvtConfigRequest         = "htmx:configRequest"
	HXEvtConfirm               = "htmx:confirm"
	HXEvtHistoryCacheError     = "htmx:historyCacheError"
	HXEvtHistoryCacheMiss      = "htmx:historyCacheMiss"
	HXEvtHistoryCacheMissError = "htmx:historyCacheMissError"
	HXEvtHistoryCacheMissLoad  = "htmx:historyCacheMissLoad"
	HXEvtHistoryRestore        = "htmx:historyRestore"
	HXEvtBeforeHistorySave     = "htmx:beforeHistorySave"
	HXEvtLoad                  = "htmx:load"
	HXEvtNoSSESourceError      = "htmx:noSSESourceError"
	HXEvtOnLoadError           = "htmx:onLoadError"
	HXEvtOobAfterSwap          = "htmx:oobAfterSwap"
	HXEvtOobBeforeSwap         = "htmx:oobBeforeSwap"
	HXEvtOobErrorNotarget      = "htmx:oobErrorNoTarget"
	HXEvtPrompt                = "htmx:prompt"
	HXEvtPushedIntoHistory     = "htmx:pushedIntoHistory"
	HXEvtResponseError         = "htmx:responseError"
	HXEvtSendAbort             = "htmx:sendAbort"
	HXEvtSendError             = "htmx:sendError"
	HXEvtSseError              = "htmx:sseError"
	HXEvtSseOpen               = "htmx:sseOpen"
	HXEvtSwapError             = "htmx:swapError"
	HXEvtTargetError           = "htmx:targetError"
	HXEvtTimeout               = "htmx:timeout"
	HXEvtValidationValidate    = "htmx:validation:validate"
	HXEvtValidationFailed      = "htmx:validation:failed"
	HXEvtValidationHalted      = "htmx:validation:halted"
	HXEvtXhrAbort              = "htmx:xhr:abort"
	HXEvtXhrLoadend            = "htmx:xhr:loadend"
	HXEvtXhrLoadstart          = "htmx:xhr:loadstart"
	HXEvtXhrProgress           = "htmx:xhr:progress"

	HXSwapInnerHTML   = "innerHTML"
	HXSwapOuterHTML   = "outerHTML"
	HXSwapAfterbegin  = "afterbegin"
	HXSwapBeforebegin = "beforebegin"
	HXSwapBeforeend   = "beforeend"
	HXSwapAfterend    = "afterend"
	HXSwapDelete      = "delete"
	HXSwapNone        = "none"
)

//
// Core Attributes
//

// HXGet creates an hx-get attribute
func HXGet(url string) Node {
	return &AttributeNode{Name: AttrHXGet, Value: url}
}

// HXPost creates an hx-post attribute
func HXPost(url string) Node {
	return &AttributeNode{Name: AttrHXPost, Value: url}
}

// HXOn creates an hx-on:event attribute
func HXOn(event, script string) Node {
	return &AttributeNode{Name: "hx-on:" + event, Value: script}
}

// HXPushUrl creates an hx-push-url attribute
func HXPushUrl(push string) Node {
	return &AttributeNode{Name: AttrHXPushUrl, Value: push}
}

// HXSelect creates an hx-select attribute
func HXSelect(target string) Node {
	return &AttributeNode{Name: AttrHXSelect, Value: target}
}

// HXSelectOob creates an hx-select-oob attribute
func HXSelectOob(target string) Node {
	return &AttributeNode{Name: AttrHXSelectOob, Value: target}
}

// HXSwap creates an hx-swap attribute
func HXSwap(target string) Node {
	return &AttributeNode{Name: AttrHXSwap, Value: target}
}

// HXSwapOob creates an hx-swap-oob attribute
func HXSwapOob(target string) Node {
	return &AttributeNode{Name: AttrHXSwapOob, Value: target}
}

// HXTarget creates an hx-target attribute
func HXTarget(target string) Node {
	return &AttributeNode{Name: AttrHXTarget, Value: target}
}

// HXTrigger creates an hx-trigger attribute
func HXTrigger(trigger string) Node {
	return &AttributeNode{Name: AttrHXTrigger, Value: trigger}
}

// HXVals creates an hx-vals attribute
func HXVals(vals string) Node {
	return &AttributeNode{Name: AttrHXVals, Value: vals}
}

//
// Additional Attributes
//

// HXBoost creates an hx-boost attribute
func HXBoost(boost ...bool) Node {
	v := ValTrue
	if len(boost) > 0 && !boost[0] {
		v = ValFalse
	}
	return &AttributeNode{Name: AttrHXBoost, Value: v}
}

// HXConfirm creates an hx-confirm attribute
func HX(message string) Node {
	return &AttributeNode{Name: AttrHXConfirm, Value: message}
}

// HXDelete creates an hx-delete attribute
func HXDelete(url string) Node {
	return &AttributeNode{Name: AttrHXDelete, Value: url}
}

// HXDisable creates an hx-disable attribute
func HXDisable() Node {
	return &AttributeNode{Name: AttrHXDisable, Value: ""}
}

// HXDisabledElt creates an hx-disabled-elt attribute
func HXDisabledElt(target string) Node {
	return &AttributeNode{Name: AttrHXDisabledElt, Value: target}
}

// HXDisinherit creates an hx-disinherit attribute
func HXDisinherit(target string) Node {
	return &AttributeNode{Name: AttrHXDisinherit, Value: target}
}

// HXEncoding creates an hx-encoding attribute
func HXEncoding(encoding string) Node {
	return &AttributeNode{Name: AttrHXEncoding, Value: encoding}
}

// HXExt creates an hx-ext attribute
func HXExt(ext string) Node {
	return &AttributeNode{Name: AttrHXExt, Value: ext}
}

// HXHeaders creates an hx-headers attribute
func HXHeaders(headers string) Node {
	return &AttributeNode{Name: AttrHXHeaders, Value: headers}
}

// HXHistory creates an hx- attribute
func HXHistory(enable bool) Node {
	v := ValTrue
	if !enable {
		v = ValFalse
	}
	return &AttributeNode{Name: AttrHXHistory, Value: v}
}

// HXHistoryElt creates an hx-history-elt attribute
func HXHistoryElt() Node {
	return &AttributeNode{Name: AttrHXHistoryElt, Value: ""}
}

// HXInclude creates an hx-include attribute
func HXInclude(selector string) Node {
	return &AttributeNode{Name: AttrHXInclude, Value: selector}
}

// HXIndicator creates an hx-indicator attribute
func HXIndicator(selector string) Node {
	return &AttributeNode{Name: AttrHXIndicator, Value: selector}
}

// HXInherit creates an hx-inherit attribute
func HXInherit(attr string) Node {
	return &AttributeNode{Name: AttrHXInherit, Value: attr}
}

// HXParams creates an hx-params attribute
func HXParams(params string) Node {
	return &AttributeNode{Name: AttrHXParams, Value: params}
}

// HXPatch creates an hx-patch attribute
func HXPatch(url string) Node {
	return &AttributeNode{Name: AttrHXPatch, Value: url}
}

// HXPreserve creates an hx-preserve attribute
func HXPreserve(preserve bool) Node {
	if !preserve {
		return nil
	}
	return &AttributeNode{Name: AttrHXPreserve, Value: ""}
}

// HXPrompt creates an hx-prompt attribute
func HXPrompt(prompt string) Node {
	return &AttributeNode{Name: AttrHXPrompt, Value: prompt}
}

// HXPut creates an hx-put attribute
func HXPut(url string) Node {
	return &AttributeNode{Name: AttrHXPut, Value: url}
}

// HXReplaceUrl creates an hx-replace-url attribute
func HXReplaceUrl(replace string) Node {
	return &AttributeNode{Name: AttrHXReplaceUrl, Value: replace}
}

// HXRequest creates an hx-request attribute
func HXRequest(config string) Node {
	return &AttributeNode{Name: AttrHXRequest, Value: config}
}

// HXSync creates an hx-sync attribute
func HXSync(selector string) Node {
	return &AttributeNode{Name: AttrHXSync, Value: selector}
}

// HXValidate creates an hx-validate attribute
func HXValidate(validate bool) Node {
	v := ValFalse
	if validate {
		v = ValTrue
	}
	return &AttributeNode{Name: AttrHXValidate, Value: v}
}
