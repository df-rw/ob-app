---
sql:
    data: ./data/rand-xy.csv
---
# observable app

This is the home page of your new Observable Framework project.

For more, see <https://observablehq.com/framework/getting-started>.

<!-- a couple of calls to the backend api -->
<div class="grid grid-cols-2">
    <div class="card">
        <h2>server time</h2>
        <button
            hx-get="/api/now"
            hx-target="#now"
            hx-swap="innerHTML">hit me</button>
        <span id="now"></span>
    </div>
    <div class="card">
        <h2>server time + 10minutes</h2>
        <button
            hx-get="/api/then"
            hx-target="#then"
            hx-swap="innerHTML">hit me</button>
        <span id="then"></span>
    </div>
</div>

<!-- making sure deployment works with sql -->
```sql id=data
SELECT * FROM data
```

<!-- making sure deployment works with file attachments -->
```js
const div2 = FileAttachment("./data/rand-xy-div2.csv").csv({ typed: true })
```

<!-- and making sure deployment works with plot -->
```js
resize(width =>
  Plot.plot({
    width,
    height: 200,
    x: {
      domain: [0, 100],
    },
    y: {
      domain: [0, 100],
    },
    marks: [
      Plot.axisX({
        ticks: d3.ticks(0, 100, 10),
      }),
      Plot.axisY({
        ticks: d3.ticks(0, 100, 10),
      }),
      Plot.dot(data, {
        x: 'x',
        y: 'y',
        stroke: "green",
      }),
      Plot.dot(div2, {
        x: 'x',
        y: 'y',
        stroke: "blue",
      }),
    ],
  }),
)
```
