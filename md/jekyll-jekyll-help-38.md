# [Jekyll 2 fails to auto-gen on folder starting with &#x27;___&#x27; outside of Jekyll scope](https://github.com/jekyll/jekyll-help/issues/38)

> state: **closed** opened by: **** on: ****

Hey guys,

A few of my projects have stopped working, and I realize it&#x27;s because my folders are set up kind of like this:
Documents/Folders/___ImportantFolder/JekyllProjects/Project1
Documents/Folders/___ImportantFolder/JekyllProjects/Project2

Ideally I want Jekyll&#x27;s escaping &#x27;___&#x27; only to happen within Project1 and Proejct2, so it&#x27;s not generating folders that start with ___ in Project1 or 2, but right now it&#x27;s stopped generating anything that starts with ___, even when those folders are outside of Jekyll&#x27;s project folder&#x27;s scope

This was never a problem with the previous Jekyll.

Thanks!

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/38#issuecomment-55624789) on: **Monday, September 15, 2014**

@janzheng Is this still a problem? I just tried recreating this in 2.4.0, but it seems to work fine.

&#x60;&#x60;&#x60;
Troys-MacBook-Air:troyswanson.github.io troy$ jekyll --version
jekyll 2.4.0
Troys-MacBook-Air:troyswanson.github.io troy$ jekyll build
Configuration file: /Users/troy/__Projects/troyswanson.github.io/_config.yml
            Source: /Users/troy/__Projects/troyswanson.github.io
       Destination: /Users/troy/__Projects/troyswanson.github.io/_site
      Generating... 
                    done.
 Auto-regeneration: disabled. Use --watch to enable.
&#x60;&#x60;&#x60;

If it&#x27;s still broken for you, feel free to re-open the issue and include some additional info about your setup, e.g.: operating system, Ruby version, etc.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/38#issuecomment-55625714) on: **Monday, September 15, 2014**

@troyswanson I&#x27;m not sure, if this is an issue, but he uses triple underscores (e.g. &#x60;___&#x60; instead of &#x60;__&#x60;).
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/38#issuecomment-55626185) on: **Monday, September 15, 2014**

Ah, you&#x27;re right. Either way, it still works.

&#x60;&#x60;&#x60;
Troys-MacBook-Air:troyswanson.github.io troy$ cd ../..
Troys-MacBook-Air:~ troy$ mv __Projects ___Projects
Troys-MacBook-Air:~ troy$ cd ___Projects/troyswanson.github.io/
Troys-MacBook-Air:troyswanson.github.io troy$ jekyll build
Configuration file: /Users/troy/___Projects/troyswanson.github.io/_config.yml
            Source: /Users/troy/___Projects/troyswanson.github.io
       Destination: /Users/troy/___Projects/troyswanson.github.io/_site
      Generating... 
                    done.
 Auto-regeneration: disabled. Use --watch to enable.
&#x60;&#x60;&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/38#issuecomment-55626439) on: **Monday, September 15, 2014**

Also, &#x60;--watch&#x60; is working fine.

&#x60;&#x60;&#x60;
Troys-MacBook-Air:troyswanson.github.io troy$ jekyll build --watch
Configuration file: /Users/troy/___Projects/troyswanson.github.io/_config.yml
            Source: /Users/troy/___Projects/troyswanson.github.io
       Destination: /Users/troy/___Projects/troyswanson.github.io/_site
      Generating... 
                    done.
 Auto-regeneration: enabled for &#x27;/Users/troy/___Projects/troyswanson.github.io&#x27;
      Regenerating: 1 files at 2014-09-15 12:30:51 ...done.
&#x60;&#x60;&#x60;
