# [Weird regeneration problem - certain folder always capitalized.](https://github.com/jekyll/jekyll-help/issues/130)

> state: **open** opened by: **** on: ****

If i have a file named videos.html - not in the _posts folder.

with this:

&#x60;&#x60;&#x60;
---
permalink: /videos/
layout: default
title: Videos
description: 
---
&lt;div class=&quot;container content&quot;&gt;
        &lt;div class=&quot;row&quot;&gt;
            &lt;div class=&quot;col-md-12&quot;&gt;
hi
            &lt;/div&gt;
        &lt;/div&gt;
&lt;/div&gt;    
&#x60;&#x60;&#x60;
no matter what I do the output folder comes out capitalized as Videos.
If I change the permalink to vvideos then it is correct as vvideos. I can rename the file something else and it doesn&#x27;t change.

I have created a new jekyll site with nothing in it other than the the stock stuff - does the same thing.

I have also rebooted twice - no change. I also changed the title and still does it. No matter what I do if I have &#x60;&#x60;&#x60;permalink: /videos/&#x60;&#x60; then I get the output folder as Videos.

I can put anything else in the permalink and it works fine. Just doesn&#x27;t like the word videos.

Is there some sort of cache for ruby or jekyll that I need to clear?

I am on a windows 8 with the latest version of jekyll 2.3.0

Help.

### Comments

---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/130#issuecomment-52521587) on: **Monday, August 18, 2014**

I just used a different computer to generate my site, and this one is generating it properly - the folder is not capitalized.

There has to be some cache somewhere that I can clear - any ideas?
This computer is Win 7, same jekyll version 2.3

I tired it again on the other computer (win 8) this morning and the folder was still capitalized.
