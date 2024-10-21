---
title: htmx events
---

```js
const logEvent = (e) => {
    console.log(e.type)
};
const events = [
    { name: 'htmx:afterOnLoad' },
    { name: 'htmx:afterProcessNode' },
    { name: 'htmx:afterRequest' },
    { name: 'htmx:afterSettle' },
    { name: 'htmx:afterSwap' },
    { name: 'htmx:beforeCleanupElement' },
    { name: 'htmx:beforeOnLoad' },
    { name: 'htmx:beforeProcessNode' },
    { name: 'htmx:beforeRequest' },
    { name: 'htmx:beforeSwap' },
    { name: 'htmx:beforeSend' },
    { name: 'htmx:beforeTransition' },
    { name: 'htmx:configRequest' },
    { name: 'htmx:confirm' },
    { name: 'htmx:historyCacheError' },
    { name: 'htmx:historyCacheMiss' },
    { name: 'htmx:historyCacheMissError' },
    { name: 'htmx:historyCacheMissLoad' },
    { name: 'htmx:historyRestore' },
    { name: 'htmx:beforeHistorySave' },
    { name: 'htmx:load' },
    { name: 'htmx:noSSESourceError' },
    { name: 'htmx:onLoadError' },
    { name: 'htmx:oobAfterSwap' },
    { name: 'htmx:oobBeforeSwap' },
    { name: 'htmx:oobErrorNoTarget' },
    { name: 'htmx:prompt' },
    { name: 'htmx:pushedIntoHistory' },
    { name: 'htmx:responseError' },
    { name: 'htmx:sendError' },
    { name: 'htmx:sseError' },
    { name: 'htmx:sseOpen' },
    { name: 'htmx:targetError' },
    { name: 'htmx:timeout' },
    { name: 'htmx:validation:validate' },
    { name: 'htmx:validation:failed' },
    { name: 'htmx:validation:halted' },
    { name: 'htmx:xhr:abort' },
    { name: 'htmx:xhr:loadend' },
    { name: 'htmx:xhr:loadstart' },
    { name: 'htmx:xhr:progress' },
]
events.forEach((e) => {
    document.addEventListener(e.name, logEvent);
});
```

<button hx-get="/api/now" hx-target="#time">get server time</button>

<p id="time"></p>
