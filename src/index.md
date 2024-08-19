# observable app

This is the home page of your new Observable Framework project.

For more, see <https://observablehq.com/framework/getting-started>.

<button
    hx-get="/api/now"
    hx-target="#now"
    hx-swap="innerHTML">now</button>

<button
    hx-get="/api/then"
    hx-target="#then"
    hx-swap="innerHTML">then</button>

<div id="now"></div>
<div id="then"></div>
