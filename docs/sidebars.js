/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

module.exports = {
  // By default, Docusaurus generates a sidebar from the docs folder structure
  // docsSidebar: [{type: 'autogenerated', dirName: '.'}],

  docsSidebar: [
    'introduction',
    {
      type: "category",
      label: "Guides",
      items: [
        "guides/overview",
        "guides/providers",
        "guides/receivers",
        "guides/subscriptions",
        "guides/rules",
        "guides/templates",
        "guides/alert_history",
        "guides/bulk_rules",
        "guides/monitoring",
        "guides/deployment",
        "guides/troubleshooting",
      ],
    },
    {
      type: "category",
      label: "Concepts",
      items: [
        "concepts/overview",
        "concepts/architecture",
        "concepts/schema",
      ],
    },
    {
      type: "category",
      label: "Contribute",
      items: ["contribute/contribution", "contribute/release"],
    },
    {
      type: "category",
      label: "Reference",
      items: ["reference/api", "reference/configuration",],
    },
  ],
};
