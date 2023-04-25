(def start
  `<html lang="en">
  <head>
    <style>
    html {
      max-width: 70ch;
      padding: 3em 1em;
      margin: auto;
      line-height: 1.75;
      font-size: 1.25em;
    }
    </style>
  </head>
  <body>
  <article>`)

(def end
  `</article>
  </body>
  </html>`)

(defn main [&]
  (def input (string/trim (file/read stdin :all)))
  (print (string start input end)))
