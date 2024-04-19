/**
 * @typedef {Object} JsDescriptor JavaScript descriptor
 * @property {string} src URL
 * @property {string} type MIME type
 * @property {string} [integrity] Subresource Integrity
 */

/**
 * @typedef {Object} CssDescriptor CSS descriptor
 * @property {string} href URL
 * @property {string} [integrity] Subresource Integrity
 */

/**
 * @typedef {Object} MetaDescriptor Meta descriptor
 * @property {string} name
 * @property {string} content
 */

/**
 * @typedef {Object} HtmlLayoutProps
 * @property {string} lang Language
 * @property {string} title Title
 * @property {JsDescriptor[]=} js JavaScript
 * @property {CssDescriptor[]=} css CSS
 * @property {MetaDescriptor[]=} meta Meta
 */

/**
 *
 * @param {HtmlLayoutProps} props Properties
 * @param {string} content HTML content
 * @returns {string} HTML layout
 */
export default function HtmlLayout({
                                       lang,
                                       title,
                                       js = [],
                                       css = [],
                                       meta = [],
                                   }, content) {
    return `
    <!DOCTYPE html>
    <html lang="${lang}">
      <head>
        <title>${title}</title>
        ${meta.map(({name, content}) =>
        `<meta name="${name}" content="${content}">`).join('\n')}
        
        ${js.map(({src, type, integrity}) =>
        `<script src="${src}" type="${type}"${integrity ? ` integrity="${integrity}"` : ''}></script>`).join('\n')}
        
        <link rel="stylesheet" href="/assets/css/primeflex.css">
        <link rel="stylesheet" href="/assets/css/primeflex-light.css">
        
        ${css.map(({href, integrity}) =>
        `<link rel="stylesheet" href="${href}"${integrity ? ` integrity="${integrity}"` : ''}>`).join('\n')}
      </head>
      <body>
        ${content}
      </body>
    </html>
  `;
}

