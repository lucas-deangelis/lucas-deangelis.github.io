# This blog

The way this blog "engine" is made is an experiment on what the laziest way to do a blog could look like. Here's what I need from a blog engine:

- content storage
- deployment
- nice appearance
- templating (all content should go into the same skeleton so I don't write it multiple times)
- markdown rendering

I'll explain how I solved everything.

## Content storage

The blog engine itself and all the posts are stored in git.

## Deployment

I'll use either github pages or netlify.

## Nice appearance

I added [sakura](https://github.com/oxalorg/sakura) in the head of the HTML skeleton. It looks nice, and also comes with a dark theme that I'll try to use with media queries.

## Templating

This is done with the following `generate.js`:

```javascript
const fs = require("fs");

let post = fs.readFileSync("_post.html").toString("utf8");

let files = fs.readdirSync("./_posts/");

for (const file in files) {
    let filename = files[file];
    let filenameWithoutMD = filename.split(".")[0];
    let content = fs.readFileSync("./_posts/" + filename).toString("utf8");
    let newPost = post.replace("# replace me with actual content #", content);

    fs.writeFileSync(filenameWithoutMD + ".html", newPost);
}
```

Not much to say about it. It works well.

## Markdown rendering

This is where it gets a bit ugly, or maybe a bit evil. The markdown is rendered on the client. Here's the HTML skeleton:

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>My blog</title>
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <link rel="stylesheet" href="https://unpkg.com/sakura.css/css/sakura.css" type="text/css">
  </head>
  <body>
    <div id="content"></div>
    <div id="md" hidden># replace me with actual content #</div>
    <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
    <script>
      document.getElementById("content").innerHTML = marked(
        document.getElementById("md").innerText
      );
    </script>
  </body>
</html>
```

It means the whole `marked.min.js` is shipped to every users, and the rendering executed in their browser. While this is a very inefficient, I wanted to see if doing something like this still produced a leaner blog that what we see on the internet. Tracking, ads, big pictures, videos can takes a long time to load.

## The end

That's pretty much it. I'll soon move the rendering into the `generate.js` script, at which point I'll have my very own static site generator.