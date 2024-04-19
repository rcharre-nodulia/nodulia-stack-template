import HtmlLayout from "../layouts/HtmlLayout.js";

/**
 * @returns {string} The Home Page
 */
export default function HomePage(){
    return HtmlLayout({
        lang: 'fr',
        title: 'Home Page',
    }, `
        <h1>Home Page</h1>
        <p>Welcome to my app!</p>
    `);
}

