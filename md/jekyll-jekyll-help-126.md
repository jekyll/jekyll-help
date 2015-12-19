# [jekyll watch not watching](https://github.com/jekyll/jekyll-help/issues/126)

> state: **closed** opened by: **** on: ****

I have a weird issue where Jekyll isn&#x27;t watching for changes when the project in certain directory. I have project in the location of &#x60;&#x60;users/me/projects/myproject/&#x60;&#x60; which if I cd into with the &#x60;&#x60;--watch&#x60;&#x60; flag works as expected. The project rebuilds on changes. 

If I move the project to my dropbox and run the same commands the project generates but no longer watches for changes.

I tested this with a new project too and the same issue is reproduced.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/126#issuecomment-52507781) on: **Monday, August 18, 2014**

Can you give some more info about your setup? Which version of Jekyll are you running, which operating system are you using, etc?
---
> from: [**qasimalyas**](https://github.com/jekyll/jekyll-help/issues/126#issuecomment-52744937) on: **Wednesday, August 20, 2014**

I&#x27;m using the current version of jekyll (2.3.0) which was upgraded from 1.5. I&#x27;m on Mac (Mavericks). My project is on github https://github.com/qasimalyas/Jekyll-site but my config is out of date there but it should be

&#x60;&#x60;&#x60;
markdown: kramdown
highlghter: pygments
permalink: /:day-:month-:year/:title.html
&#x60;&#x60;&#x60;
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/126#issuecomment-52746201) on: **Wednesday, August 20, 2014**

&#x60;jekyll watch&#x60; does not register changes to &#x60;_config.yml&#x60;; you have to restart it.

We probably need to communicate this in a better way with the CLI.

Just so that’s clear; it’s not a given that it resolves your problem.
---
> from: [**qasimalyas**](https://github.com/jekyll/jekyll-help/issues/126#issuecomment-52750148) on: **Wednesday, August 20, 2014**

I&#x27;m not at my Mac right now but the problem was also reproduced if I created a [new Jekyll project](http://jekyllrb.com/docs/quickstart/). Again this is only if I&#x27;m in my Dropbox folder. As soon as I moved the project out &#x60;&#x60;--watch&#x60;&#x60; works as expected.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/126#issuecomment-52791529) on: **Wednesday, August 20, 2014**

My guess is that Dropbox is conflicting with Jekyll&#x27;s watcher. Dropbox&#x27;s app might be causing other watchers to break. Try killing the Dropbox app and see how it goes.
---
> from: [**qasimalyas**](https://github.com/jekyll/jekyll-help/issues/126#issuecomment-52883000) on: **Wednesday, August 20, 2014**

Seems to be working now as normal. Bizarre. Closing issue.
