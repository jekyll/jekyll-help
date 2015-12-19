# [Problem with permalink case and existing folder names on the same drive - Windows](https://github.com/jekyll/jekyll-help/issues/250)

> state: **open** opened by: **** on: ****

I had an issue where no matter what I did, if I had &#x60;&#x60;&#x60;permalink: /videos/&#x60;&#x60;&#x60; I would get Videos (uppercase V) as the folder name.

It turned out that I had a folder on the same drive but outside the jekyll project that was also called Videos with a cap V. That folder was a parent folder on the same drive. It looks like any parent/top level folder on the same drive that has the same name as the permalink causes jekyll to pickup the case of that folder. So if I change the folder to &#x27;videos&#x27; then jekyll gives me &#x27;videos&#x27; and all is well. 

So I had a folder named Videos (outside of my jekyll project), jekyll gave me Videos when I asked for videos.
If I have a folder named TEST-images, then I get TEST-images when I ask for test-images in the permalink.

It doesn&#x27;t appear to do it on folders that are not top level/parent folders. So if the folder &#x27;TEST-images&#x27; is inside another folder named &#x27;test&#x27; then there is no problem. It also doesn&#x27;t seem to be a problem if the folders are on another drive - it has to be same drive, parent folder.

Just to be clear, the folders causing the problem are outside of jekyll, and have nothing to do with jekyll, they are just random other folders on the same drive, but if by chance you want a permalink that is the same name as one of these other folder names, jekyll picks up the case of the outside folder rather than giving you what you asked for. I suppose I should be calling them directories?

Windows 8.1, latest Jekyll - 2.5.3

### Comments

