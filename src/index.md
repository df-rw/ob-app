# observable app

This is the home page of your new Observable Framework project.

For more, see <https://observablehq.com/framework/getting-started>.

The appserver path is: ${APPSERVER}

```js
const makeGetNow = (label) => {
    return html`<button
        hx-get="${APPSERVER}/now"
        hx-target="#now"
        hx-swap="innerHTML"
    >${label}</button>`;
};
```

```js
display(makeGetNow("get now"));
await visibility()
htmx.process(document.body);
```

<div id="now"></div>
