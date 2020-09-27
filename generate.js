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