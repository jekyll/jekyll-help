# [push to github](https://github.com/jekyll/jekyll-help/issues/224)

> state: **open** opened by: **** on: ****

I have my ragwing.github.io working, but I cannot push the new changes.
I had it in the wrong repo, so copied it to the one above and it worked. 
Just cannot push now.

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/224#issuecomment-68609144) on: **Saturday, January 03, 2015**

That doesn&#x27;t sound like a Jekyll issue at all. What are you typing into the command line after changing some of your files?
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/224#issuecomment-68633768) on: **Sunday, January 04, 2015**

I did a git add remote ragwing ragwing.github.io
then did push ragwing

On Sat, Jan 3, 2015 at 3:52 PM, Philipp Rudloff &lt;notifications@github.com&gt;
wrote:

&gt; That doesn&#x27;t sound like a Jekyll issue at all. What are you typing into
&gt; the command line after changing some of your files?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub.
&gt;



-- 
Rex Guinn
Pictures http://picasaweb.google.com/ragwing
My SIte http:// &lt;http://rexguinn.bluegrassareawoodturners.org&gt;rexguinn.net
BAW http://bluegrassareawoodturners.org
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/224#issuecomment-68659769) on: **Sunday, January 04, 2015**

Do &#x60;git fetch&#x60; and then &#x60;git status&#x60; to see how many commits ahead your clone of the repo is.

Also, what kind of error is it giving you?Just saying you can&#x27;t do something doesn&#x27;t help us.
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/224#issuecomment-68809055) on: **Monday, January 05, 2015**

git fetch  ragwing
fatal: Couldn&#x27;t find remote ref ragwing
fatal: The remote end hung up unexpectedly
guinn@laptop:~/ragwing-master$ git push . ragwing
error: src refspec ragwing does not match any.
error: failed to push some refs to &#x27;.&#x27;
guinn@laptop:~/ragwing-master$ git push all ragwing
error: src refspec ragwing does not match any.
error: failed to push some refs to &#x27;all&#x27;
guinn@laptop:~/ragwing-master$ git fetch ragwing
fatal: &#x27;ragwing.github.io&#x27; does not appear to be a git repository
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.

git status
On branch gh-pages

Initial commit

Changes to be committed:
  (use &quot;git rm --cached &lt;file&gt;...&quot; to unstage)

new file:   README.markdown
new file:   _deploy/README.markdown
new file:   _deploy/about.html
new file:   _deploy/blog.html
new file:   _deploy/custom.css
new file:   _deploy/images/VLA.jpg
new file:   _deploy/images/almo.jpg
new file:   _deploy/images/ferry_powell.jpg
new file:   _deploy/images/flowers_tx_roadside.jpg
new file:   _deploy/images/fredericksburg.jpg
new file:   _deploy/images/hwy-78.jpg
new file:   _deploy/images/jaraffes.jpg
new file:   _deploy/images/leaving_lex.jpg
new file:   _deploy/images/lebanon.jpg
new file:   _deploy/images/logo.gif
new file:   _deploy/images/lukenbach.jpg
new file:   _deploy/images/n-of-abq.jpg
new file:   _deploy/images/nation_military_park.jpg
new file:   _deploy/images/ok_city.jpg
new file:   _deploy/images/padre.jpg
new file:   _deploy/images/rex.jpg
new file:   _deploy/images/sadeona.jpg
new file:   _deploy/images/st_pet_fl.jpg
new file:   _deploy/images/subscribe-icon.gif
new file:   _deploy/images/subscribe.png
new file:   _deploy/images/tucson.jpg
new file:   _deploy/images/usa.gif
new file:   _deploy/index.html
new file:   _deploy/javascripts/jquery.github.js
new file:   _deploy/javascripts/jquery.js
new file:   _deploy/open-source.html
new file:   _deploy/posts/Leaving-Lexington/index.html
new file:   _deploy/posts/Welcome-to-My-Cool-Site/index.html
new file:   _deploy/posts/after-keywest/index.html
new file:   _deploy/posts/to-california/index.html
new file:   _deploy/posts/vicksburg-to-padre/index.html
new file:   _deploy/railsgit/index.html
new file:   _deploy/resetvcheckout.html
new file:   _deploy/stylesheets/master.css
new file:   _layouts/master.html
new file:   _layouts/post.html
new file:   _posts/2002-03-07-Leaving-Lexington.md
new file:   _posts/2002-03-20-after-keywest.md
new file:   _posts/2002-5-20-to-california.md
new file:   _posts/links.md
new file:   _site/README.markdown
new file:   _site/about.html
new file:   _site/blog.html
new file:   _site/custom.css
new file:   _site/images/VLA.jpg
new file:   _site/images/almo.jpg
new file:   _site/images/ferry_powell.jpg
new file:   _site/images/flowers_tx_roadside.jpg
new file:   _site/images/fredericksburg.jpg
new file:   _site/images/hwy-78.jpg
new file:   _site/images/jaraffes.jpg
new file:   _site/images/leaving_lex.jpg
new file:   _site/images/lebanon.jpg
new file:   _site/images/logo.gif
new file:   _site/images/lukenbach.jpg
new file:   _site/images/n-of-abq.jpg
new file:   _site/images/nation_military_park.jpg
new file:   _site/images/ok_city.jpg
new file:   _site/images/padre.jpg
new file:   _site/images/rex.jpg
new file:   _site/images/sadeona.jpg
new file:   _site/images/st_pet_fl.jpg
new file:   _site/images/subscribe-icon.gif
new file:   _site/images/subscribe.png
new file:   _site/images/tucson.jpg
new file:   _site/images/usa.gif
new file:   _site/index.html
new file:   _site/javascripts/jquery.github.js
new file:   _site/javascripts/jquery.js
new file:   _site/open-source.html
new file:   _site/resetvcheckout.html
new file:   _site/stylesheets/master.css
new file:   about.textile
new file:   blog.html
new file:   custom.css
new file:   images/VLA.jpg
new file:   images/almo.jpg
new file:   images/ferry_powell.jpg
new file:   images/flowers_tx_roadside.jpg
new file:   images/fredericksburg.jpg
new file:   images/hwy-78.jpg
new file:   images/jaraffes.jpg
new file:   images/leaving_lex.jpg
new file:   images/lebanon.jpg
new file:   images/logo.gif
new file:   images/lukenbach.jpg
new file:   images/n-of-abq.jpg
new file:   images/nation_military_park.jpg
new file:   images/ok_city.jpg
new file:   images/padre.jpg
new file:   images/rex.jpg
new file:   images/sadeona.jpg
new file:   images/st_pet_fl.jpg
new file:   images/subscribe-icon.gif
new file:   images/subscribe.png
new file:   images/tucson.jpg
new file:   images/usa.gif
new file:   index.html
new file:   index1.html
new file:   javascripts/jquery.github.js
new file:   javascripts/jquery.js
new file:   open-source.html
new file:   resetvcheckout.html
new file:   rex-blog
new file:   stylesheets/master.css


On Sun, Jan 4, 2015 at 8:45 PM, Troy Swanson &lt;notifications@github.com&gt;
wrote:

&gt; Do git fetch and then git status to see how many commits ahead your clone
&gt; of the repo is.
&gt;
&gt; Also, what kind of error is it giving you?Just saying you can&#x27;t do
&gt; something doesn&#x27;t help us.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/224#issuecomment-68659769&gt;.
&gt;



-- 
Rex Guinn
Pictures http://picasaweb.google.com/ragwing
My SIte http:// &lt;http://rexguinn.bluegrassareawoodturners.org&gt;rexguinn.net
BAW http://bluegrassareawoodturners.org
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/224#issuecomment-68936230) on: **Tuesday, January 06, 2015**

If I were you, I would archive the folder that has your changes in it and re-clone the repo. Then, check out the &#x60;gh-pages&#x60; branch and copy the files from the archived folder into the new clone.
