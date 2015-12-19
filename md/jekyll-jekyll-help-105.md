# [Jekyll site in subdirectory on Pages – Working with site.url locally?](https://github.com/jekyll/jekyll-help/issues/105)

> state: **closed** opened by: **** on: ****

When powering a site with Jekyll from a subdirectory with GitHub Pages, one prepends &#x60;{{ site.url }}&#x60; (or something similar) to internal links/assets etc.

Now when starting a server locally, all assets are pulled from the actual site specified in the _config.yml. How does one develop locally without always changing the configuration?

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-50651327) on: **Wednesday, July 30, 2014**

I use a script/preview file which uses a custom dev config that is passed along with the main config file. i use that locally and pages runs like normal.
---
> from: [**ivantsepp**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-50658061) on: **Wednesday, July 30, 2014**

Would it be a good idea if we could set &quot;development&quot; configuration in the config file? One would set in &#x60;_config.yml&#x60;:
&#x60;&#x60;&#x60;yaml
development:
  url: &quot;localhost:4000&quot;
&#x60;&#x60;&#x60;
And that configuration would override if &#x60;Jekyll.env == &#x27;development&#x27;&#x60;?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-50658666) on: **Wednesday, July 30, 2014**

@ivantsepp That&#x27;s a little more complicated than I&#x27;d prefer... just worried about confusion. I&#x27;d rather see a convention like for dotenv where &#x60;.env.development&#x60; is picked up if it exists, but in our case, it&#x27;d be &#x60;_config.#{Jekyll.env}.yml&#x60;.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-50658697) on: **Wednesday, July 30, 2014**

/cc @benbalter for his :thought_balloon: 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-50659940) on: **Wednesday, July 30, 2014**

@parkr A _config.dev.yml would duplicate the contents of the _config.yml then? Or would it override settings from the _config.yml, such as url?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-50660698) on: **Wednesday, July 30, 2014**

It would override.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-54944180) on: **Tuesday, September 09, 2014**

In the meantime the solution is &#x60;baseurl: /repo&#x60; and prepending with &#x60;{{ site.baseurl }}/img/icon.png&#x60;, etc.? Any other suggestions?

I&#x27;d love to have build tools do some magic for me there, however I never found a way to get Grunt/Gulp to work reliably on Windows.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/105#issuecomment-55643650) on: **Monday, September 15, 2014**

@kleinfreund You might want to try using [GitHub Pages Metadata](https://help.github.com/articles/repository-metadata-on-github-pages), if you&#x27;re using GitHub Pages, that is.
