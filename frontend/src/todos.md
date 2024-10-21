---
theme: dashboard
title: Todo application with htmx
---

# Todo application with htmx

```js
const makeLoadPage = (params) => {
    return html`<div hx-get="/api/todos${params}" hx-trigger="load"><p>I'm loading here</p></div>`;
};
display(makeLoadPage(window.location.search))
await visibility();
htmx.process(document.body)
```
