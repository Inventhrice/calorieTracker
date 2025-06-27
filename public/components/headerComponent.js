export default {
    props: {
        title: {
            type: String,
            required: true,
            default: "Default"
        }
    },
    template: `<div id="header" class="top-navigation">
                <h5 class="title-text">{{title}}</h5>
            </div>`
}