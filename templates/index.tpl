<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.title}}</title>
    <meta name="keywords" content="{{.keywords}}"/>
    <meta name="description" content="{{.description}}"/>

    <link type="text/css" rel="stylesheet" media="screen" href="/static/plugins/prettify.css"/>
</head>
<body>
<section id="content">
    <span>
        <textarea id="result" style="display: none">{{.content}}</textarea>
    </span>
</section>
</body>
<script type="text/javascript" language="JavaScript" src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
<script type="text/javascript" language="JavaScript" src="/static/plugins/marked.min.js"></script>
<script type="text/javascript" language="JavaScript" src="/static/plugins/prettify.js"></script>
</html>
<script type="text/javascript">
    var tmp = marked($("#result").val());
    // console.log(tmp);
    $("#content").html(tmp);

    $(document).ready(function () {
        $("pre").addClass("prettyprint linenums:0");
        prettyPrint();
    });
</script>