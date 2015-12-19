# [posts name in Hebrew](https://github.com/jekyll/jekyll-help/issues/111)

> state: **closed** opened by: **** on: ****

i made a bunch of dummy posts named in Hebrew for testing purposes, and Jekyll got stuck regenerating the site without stopping. is there a fix or should i give up on jekyll? 

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/111#issuecomment-50967104) on: **Saturday, August 02, 2014**

Well, you could provide some more information (i.e. example posts) so we can try on our own. 
---
> from: [**ghost**](https://github.com/jekyll/jekyll-help/issues/111#issuecomment-50968174) on: **Saturday, August 02, 2014**

hi i&#x27;m pretty new at this so please bear with me - 

i&#x27;m attaching screenshots of the console and the posts list, thank you

![posts](https://cloud.githubusercontent.com/assets/8005415/3788443/385cb5ea-1a66-11e4-8fdc-a27df84b095a.JPG)
![screenshot](https://cloud.githubusercontent.com/assets/8005415/3788444/3df04cb0-1a66-11e4-8dc5-ed0cc401f74b.JPG)




---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/111#issuecomment-52431585) on: **Sunday, August 17, 2014**

@isaacfeldman The file names should contain ascii only. Put the Hebrew title in the front matter. Example:

1. Create file &#x60;_posts/2014-08-17-hebrew.md&#x60;.
2. Open it and set the title.

    &#x60;&#x60;&#x60;
---
title: &quot;הגלגול הזה הכתיב במידה&quot;
---
&#x60;&#x60;&#x60;
---
> from: [**ghost**](https://github.com/jekyll/jekyll-help/issues/111#issuecomment-52454878) on: **Sunday, August 17, 2014**

oh that&#x27;s great, thank you! i&#x27;ll let you know if it works
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/111#issuecomment-55644023) on: **Monday, September 15, 2014**

@isaacfeldman Any update you can give us on this? I&#x27;m going to close out this issue for now, but feel free to re-open it if you still need help!
