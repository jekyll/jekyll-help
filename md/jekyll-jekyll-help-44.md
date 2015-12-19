# [Work with *.Rmd files and RStudio](https://github.com/jekyll/jekyll-help/issues/44)

> state: **closed** opened by: **** on: ****

Hello,

so far I was able to set up jekyll on github pages: http://magnusmetz.github.io/

I want to use it as a substitute for RPubs and publish posts written as *.Rmd files.

It seems, there are two major issues left, which were impossible for me to solve so far and I was not able to solve the problem with the information I found on the internet.

1. I use RStudio to version control the website. If I make a change in one of the markdown files in the _posts folder and commit these changes, the changes are not overtaken on http://magnusmetz.github.io/, although I can see them on http://localhost:4000/. Why is that? It seems that changes in the markdown files are tracked but the subsequent changes in _site are not transmitted to github. How to solve this issue?

2. I want to use *.Rmd files as the root files for my posts instead of markdown. How can I accomplish that? I know the blog of http://jfisher-usgs.github.io/ but I was not really able to accomplish that. But this is a pretty good example what I want to do.

Thanks in advance for further help.

Magnus

### Comments

---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42824267) on: **Monday, May 12, 2014**

By the way: Everytime I commit and push a change in one of the markdown files in _posts I reveive an Email from github with following message:

The page build failed with the following error:

A file was included in &#x60;_layouts/post.html&#x60; that is a symlink or does not exist in your &#x60;_includes&#x60; directory. For more information, see https://help.github.com/articles/page-build-failed-file-is-a-symlink.
If you have any questions please contact us at https://github.com/contact.

Maybe this helps for your answer.

Thanks.
---
> from: [**iamgabeortiz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42824581) on: **Monday, May 12, 2014**

Regarding your &#x60;&#x60;&#x60;#1&#x60;&#x60;&#x60; above RStudio is likely only tracking files for you locally.  You would need to then push your repo changes to Github so that they display on the website.  I can&#x27;t comment to your &#x60;&#x60;&#x60;#2&#x60;&#x60;&#x60;. :smiley:
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42825372) on: **Monday, May 12, 2014**

Well I am pushing the changes to Github with RStudio, but still can&#x27;t see the changes on http://magnusmetz.github.io/....
---
> from: [**iamgabeortiz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42826123) on: **Monday, May 12, 2014**

LOL.  I didn&#x27;t even know you can push changes to Github with RStudio.  Just started working with it.  You might want to take your &#x60;&#x60;&#x60;#1&#x60;&#x60;&#x60; then to RStudio on Github and post the issue there.  Especially if Jekyll is building locally and you can see it locally.  If I have any free time this week, I&#x27;ll probably trying messing with RStudio and pushing to Github.  Would be ideal with working some of my R projects. :thought_balloon: Thanks for the idea.
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42826484) on: **Monday, May 12, 2014**

No problem. It&#x27;s usually working perfectly. This is the first time I have a problem like this. The strange thing is, it doesn&#x27;t seem that RStudio is the problem. Also if I use following commands the changes in _site are not on Github:

~~~
git add -A
git commit -m &quot;Commit message&quot;
git push
~~~

---
> from: [**iamgabeortiz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42826863) on: **Monday, May 12, 2014**

Do you push your Jekyll working directory to Github and let Github build it or do you build locally and just push everything in &#x60;&#x60;&#x60;_site/&#x60;&#x60;&#x60; folder?
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42827724) on: **Monday, May 12, 2014**

Well I push the whole repository &#x60;magnusmetz.github.com&#x60;. I know jekyll since yesterday. If I build locally with
~~~
jekyll serve
~~~ 
I can see the changes immediately on http://localhost:4000/. But if I commit and push the repository &#x60;magnusmetz.github.com&#x60; jekyll-generated changes in &#x60;_site/&#x60; , which is inside &#x60;magnusmetz.github.com&#x60; are not regarded.
I really don&#x27;t understand that...
---
> from: [**iamgabeortiz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-42830369) on: **Monday, May 12, 2014**

Yea so the **local view** with &#x60;&#x60;&#x60;Jekyll -serve&#x60;&#x60;&#x60; just serves up your &#x60;&#x60;&#x60;_site&#x60;&#x60;&#x60; folder to your local machine.  You can upload everything in that folder to a webhost and your site would look exactly the same as it does locally.  With Github Pages, you are providing Github the pre-build files and allowing Github to build and serve it for you.  Is RStudio updating posts in &#x60;&#x60;&#x60;magnusmetz.github.com/_posts/&#x60;&#x60;&#x60; or &#x60;&#x60;&#x60;magnusmetz.github.com/_site/_posts/&#x60;&#x60;&#x60;?
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/44#issuecomment-43015096) on: **Tuesday, May 13, 2014**

I deleted the whole repository and made a fresh setup. Now it works. Still don&#x27;t know the answer.
