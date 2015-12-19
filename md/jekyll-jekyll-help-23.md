# [Make a theme installable and upgradeable](https://github.com/jekyll/jekyll-help/issues/23)

> state: **closed** opened by: **** on: ****

Let me first explain, how I define a *theme*. It’s a bunch of files and configurations, that a user can try out inside an existing installation *without* affecting the previous theme.

I end up in the following namespace scenario for the theme Mytheme:

* Layouts live in the directory &#x60;_mytheme&#x60;. The &#x60;_config.yml&#x60; must be changed to:

    &#x60;&#x60;&#x60;yml
layouts: ./_mytheme
    &#x60;&#x60;&#x60;

* Includes live in the directory &#x60;_includes/mytheme&#x60;.
* External libraries live in the directory &#x60;mytheme&#x60;.
* Settings live in the variable &#x60;mytheme&#x60; inside of &#x60;_config.yml&#x60;.

I provide three scattered directories. The user should never modify them by hand. She should only set the options in the config. The installation should be foolproof. The upgrade should be *one action* only. 

My ideas:

1. Maintain one theme repository. Provide a bash script that clones (git) my theme repository in a temporary directory. Then the directories get copied to the right places inside of user’s source. Existing &#x60;mytheme&#x60; files get overridden. The user executes the same script if she wants to upgrade the theme.

1. Maintain three theme repositories, one for every directory. Provide a bash script that creates three git submodules inside of user’s source. Provide a second script for the upgrade procedure.

1. Maintain three theme repositories. Use a package manager like Bower to handle installation and upgrade procedures.

What is the easiest way to install such a theme composed of three directories? How do you create installable and upgradeable themes?


### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/23#issuecomment-41449672) on: **Friday, April 25, 2014**

We&#x27;re going to solve this with Octopress Ink, by @imathis. I&#x27;m hoping to integrate Octopress Ink into Jekyll proper once the kinks have been worked out and with @imathis&#x27;s support/permission. Please take a look at this: https://github.com/octopress/ink
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/23#issuecomment-41462129) on: **Saturday, April 26, 2014**

@parkr Thank you for the hint! As I see in the [commit history](https://github.com/octopress/ink/commits/master), it will take a while until theme authors can try Octopress Ink out. I will wait 3&amp;nbsp;months and ask again.
