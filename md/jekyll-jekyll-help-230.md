# [Why should I use git clone when deploying with git post-receive hook?](https://github.com/jekyll/jekyll-help/issues/230)

> state: **open** opened by: **** on: ****

In the deployment guide (http://jekyllrb.com/docs/deployment-methods/), there is the git post-receive hook method explained. I set up my deployment system almost the same way before finding the guide. The only difference is that I don&#x27;t git clone into a temp directory, like it&#x27;s done in the guide. Instead I use git reset to write out the HEAD.

Git dir: /srv/git/zuihitsu/zuihitsu.git (a bare repository)
Worktree: /srv/git/zuihitsu/zuihitsu.worktree
Public dir: /srv/www/zuihitsu

I just:

    #!/bin/sh
    git --work-tree=/srv/git/zuihitsu/zuihitsu.worktree reset --hard master
    jekyll build --source /srv/git/zuihitsu/zuihitsu.worktree --destination /srv/www/zuihitsu

That seems to work fine. What are the pros and cons of this reseting method, versus the cloning one? This way git doesn&#x27;t have the clone the whole repo, just write out this specific branch. But is there any hazards that I&#x27;m overlooking?

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/230#issuecomment-69060337) on: **Wednesday, January 07, 2015**

@golddranks The goal of the guide was most likely to be as simple as possible and effectively deal with force pushes that modify the history of the repo. If you&#x27;re doing a hard reset, you&#x27;re probably fine. I&#x27;ll leave this issue open for a bit for others to chime in, but I don&#x27;t see a problem with your method.
