# [permalink: /videos/ always comes up capitalized](https://github.com/jekyll/jekyll-help/issues/244)

> state: **closed** opened by: **** on: ****

This is crazy, I highly doubt it is jekylls fault, but for some reason if I have a page and put
&#x60;permalink: /videos/&#x60; in the front matter, the resulting page path is /Videos/ with a capital V. If I do anything else it gives me whatever I put in (lowercase) - vvideos works, /videos/videos/ works, but if it is &#x60;permalink: /videos/&#x60; then it is capitalized. This breaks some of my urls as AWS S3 is case sensitive. I have hundreds of other pages, and they all work fine.

I have tested this with a different project, and it does the same thing. So it is not just something to do with the one file/folder, and only if the word is &#x27;videos&#x27;.

I am on the latest jekyll - 2.5.3 on a PC win 8.1.

Is there some cache somewhere or something that is stuck with Videos that I could clear? This has been happening for a few months, but I stopped working on it for a while and am now back working on it.

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/244#issuecomment-70968725) on: **Wednesday, January 21, 2015**

Are any _plugins/ in use?
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/244#issuecomment-71382627) on: **Sunday, January 25, 2015**

nope, don&#x27;t even have a _plugins folder. Also, this is on a Page, not a post just to be clear. I also have checked the config file, and nothing is wrong there. I really think it is some sort of cache problem, though this has been going on for months.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/244#issuecomment-71385430) on: **Sunday, January 25, 2015**

wow! I got frustrated and decided to make a video of it to link to here so that you don&#x27;t think I am crazy... while producing the video I noticed that I had a folder named Videos on the same drive, and thought that was a weird coincidence so I renamed that folder to videos (lowercase) and now it works correctly in Jekyll. Then I looked at my drive and saw another folder capitalized, tried that name in the permalink (lowercase) and it came out capitalized.

I tried a sub folder that was capitalized, and that one worked ok.

So it looks like any top level folder that is on the same drive could cause a problem with permalinks and capitalization.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/244#issuecomment-71386002) on: **Sunday, January 25, 2015**

I&#x27;m going to close this and create a new issue with the problem.
