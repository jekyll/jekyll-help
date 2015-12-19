# [Jekyll::Convertible#write does not get invoked on pages for Collections](https://github.com/jekyll/jekyll-help/issues/197)

> state: **closed** opened by: **** on: ****

I am on Jekyll 2.5.1
Using the Jekyll master site as an example,
if you create a simple plugin to modify the output of a Convertible,
all pages will get affected, but Collections pages will not.

I understand Collections are preliminary, but they are being used in the master Jekyll site: http://goo.gl/bGqs9t


You can verify this by throwing the following plugin into _plugins under the master Jekyll site.

module Jekyll
  module Convertible
    def write(dest)
      self.output = &#x27;foo&#x27;
      path = destination(dest)
      FileUtils.mkdir_p(File.dirname(path))
      File.open(path, &#x27;wb&#x27;) do |f|
        f.write(output)
      end
    end
  end
end

All .html and .md files will get turned into &#x27;foo&#x27;, but the /docs/ folder will not be affected by this
plugin.




### Comments

---
> from: [**camkego**](https://github.com/jekyll/jekyll-help/issues/197#issuecomment-63505137) on: **Tuesday, November 18, 2014**

There is $100 bounty on this issue: http://goo.gl/gk4Zhu
(on the bountysource.com site)
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/197#issuecomment-63537464) on: **Tuesday, November 18, 2014**

@camkego This issue would probably be better served in the main jekyll repo, as it pretains to the functionality of jekyll itself. Closing with the hope that you&#x27;ll add an issue over there :wink: 
