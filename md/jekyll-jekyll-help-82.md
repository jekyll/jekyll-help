# [[question] Example of Jekyll site using bower and bootstrap-sass?](https://github.com/jekyll/jekyll-help/issues/82)

> state: **closed** opened by: **** on: ****

I&#x27;m looking for example Jekyll sites that use (official) bootstrap-sass, preferably managed with bower. I figure there must be some idiomatic ways of using these kind of external dependencies with Jekyll.

Thanks!

Stu   

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-46864692) on: **Monday, June 23, 2014**

There&#x27;s nothing inherently special about using Bower or other external dependencies in a Jekyll site. It does make hosting with GitHub Pages a bit more complicated, since you&#x27;d have to check in the stuff that Bower pulls down into (at least the) &#x60;gh-pages&#x60; branch.

Are there any specific questions you have about using Bower or others with Jekyll? We&#x27;d love to help!
---
> from: [**ornamentist**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-46904544) on: **Monday, June 23, 2014**

Thanks for responding--much appreciated. I can see there&#x27;s a world of very useful external dependencies for Jekyll sites and I&#x27;m wondering whether the Jekyll community has converged at all on best practices for managing these dependencies?

For example, use of Bower .vs. direct inclusion of frameworks like Bootstrap and pragmatic techniques for organizing Jekyll Git repositories with multiple external dependencies. There may be exemplar Jekyll sites that do this kind of thing well?

Thanks!

Stu


---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-46905659) on: **Monday, June 23, 2014**

There aren&#x27;t any &quot;best practices&quot; that I know of that are specific to Jekyll. I guess the advice I would give you is to use the best practices for each tool you&#x27;re inclined to use. For things like Bower, you&#x27;ll have to make choices about how you want to deploy your site with regards to the dependencies. Using Grunt or Gulp to build/compress/move source files would be a good place to start, but it would definitely be tricky with GitHub Pages.

All of that said, it&#x27;s certainly possible, and would probably be a fun proof-of-concept project.

I can envision a system that includes a Jenkins rig for continuous delivery. Include scripts for Jenkins that deploy run tests and then deploys a subtree to the &#x60;gh-pages&#x60; branch or something like that.
---
> from: [**ornamentist**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-47056967) on: **Tuesday, June 24, 2014**

FWIW: The Jekyll assets plugin (https://github.com/ixti/jekyll-assets) does have some pragmatic practices for managing external dependencies (including Bootstrap)--at the cost of not being able to use github pages.

---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-47057350) on: **Tuesday, June 24, 2014**

Here is one: https://github.com/AJAlabs/jekyllkb
---
> from: [**ornamentist**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-47057734) on: **Tuesday, June 24, 2014**

Very nice--thanks.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/82#issuecomment-47247368) on: **Thursday, June 26, 2014**

@ornamentist I&#x27;m going to close this issue. If you need more help, feel free to create a new issue!
