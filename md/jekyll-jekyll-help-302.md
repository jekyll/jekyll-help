# [Custom Liquid Plugin for custom Syntax Highlighting - embedded code confuses Kramdown](https://github.com/jekyll/jekyll-help/issues/302)

> state: **open** opened by: **** on: ****

I have written this plugin:

&#x60;&#x60;&#x60;ruby
module Jekyll
    module Tags
        class Prism &lt; Liquid::Block
            def initialize(tag_name, text, tokens)
                @arg = text.strip
                super
            end

            def render(context)
                output = super(context)
                &quot;&lt;pre&gt;&lt;code class=\&quot;language-#{@arg}\&quot;&gt;#{output}&lt;/code&gt;&lt;/pre&gt;&quot;
            end
        end
    end
end 

Liquid::Template.register_tag(&#x27;prism&#x27;, Jekyll::Tags::Prism)
&#x60;&#x60;&#x60;

This is how I use it:

&#x60;&#x60;&#x60;c++
{% prism cpp %}
#include &lt;iostream&gt;

// Hello World
int main()
{
    cout &lt;&lt; &quot;hello world&quot; &lt;&lt; endl;
    int a = 10;
}
{% endprism %}
&#x60;&#x60;&#x60;

Now, the problem is, that I mostly use C++ Code on my website. When I now generate this markdown with Jekyll, then all text after &#x60;&#x60;&#x60;{% endprism %}&#x60;&#x60;&#x60; would still be included within the &#x60;&#x60;&#x60;&lt;pre&gt;&#x60;&#x60;&#x60; Tag, because Kramdown is getting confused by &#x60;&#x60;&#x60;&lt;iostream&gt;&#x60;&#x60;&#x60; If I escape it (&#x60;&#x60;&#x60;\&lt;iostream\&gt;&#x60;&#x60;&#x60;), then my plugin works as expected, but my Javascript Highlighter is getting confused.

How can I solve this situation without enabling Jekyll&#x27;s highlighter because this one does not work properly for me. I am on Windows and barely got Jekyll working after lots of certificate problems with gem. I am also porting an existing website which used prism.js for highlighting before and I find it kind of neat.

Another good help would be if someone can tell me how if there is a Ruby function that I can use to convert all &#x60;&#x60;&#x60;&lt;&gt;&#x60;&#x60;&#x60; to html special characters within my plugin. Is there something like htmlspecialchars from PHP? So I can have the original code in my markdown file. I am sorry for asking, I am really new to ruby and jekyll.

edit: I solved it by using CGI.escapeHTML(), but I still would like to know if there is a Jekyll way to solve it, just for learning.

### Comments

