# [error](https://github.com/jekyll/jekyll-help/issues/210)

> state: **closed** opened by: **** on: ****

run
 gem install jekyll
error
While executing gem ... (Gem::FilePermissionError)
    You don&#x27;t have write permissions for the /Library/Ruby/Gems/2.0.0 directory.


### Comments

---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/210#issuecomment-66935510) on: **Sunday, December 14, 2014**

As the error states you don&#x27;t have write permissions to where the gems are installed, you probably have to run the command as &#x60;&#x60;&#x60;sudo&#x60;&#x60;&#x60; although I would advice you to setup Ruby in a way where that won&#x27;t happen :)
---
> from: [**swtlovewtt**](https://github.com/jekyll/jekyll-help/issues/210#issuecomment-66942929) on: **Sunday, December 14, 2014**

thank you.
sudo ,done
