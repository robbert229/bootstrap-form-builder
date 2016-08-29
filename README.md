#bootstrap-form-builder
This is a utility I use to build more lengthy and complicated forms. It takes a json formatted input and outputs bootstrap form inputs.

##Example

    bootstrap-form-builder.exe html .\bootstrap-form-builder\application.json .\bootstrap-form-builder\application.html 


###Application.json

    {
        "fields":[
            {
                "name":"first_name",
                "type":"text",
                "label":"Enter Your First Name:,
                "placeholder":"fname",
                "required":true
            }
        ]
    }

###Application.html

    <div class="form-group">
        <label for="first_name">Enter Your First Name *</label>
        <input type="text" class="form-control" placeholder="fname" required>
    </div>