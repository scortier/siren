import React from 'react';
import Layout from '@theme/Layout';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Container from '../core/Container';
import GridBlock from '../core/GridBlock';

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Help`}
      description="Universal data observaility toolkit">
      <main>
        <Container className="textSection wrapper" background="light">
          <h1>Need help?</h1>
          <p>
            Need a bit of help? We're here for you. Check out our current issues, GitHub discussions, or get support through Slack.
          </p>
          <GridBlock
            layout="threeColumn"
            contents={[
              {
                title: 'Slack',
                content: (
                  <div>
                    The Siren team has an open source slack workspace to discuss development and support.
                    Most of the Siren discussions happen in #siren channel.
                    <br /><a href="https://odpf-community.slack.com/"> Join us on Slack </a>
                  </div>)
              },
              {
                title: 'GitHub Issues',
                content: (
                  <div>
                    Have a general issue or bug that you've found? We'd love to hear about it in our GitHub issues. This can be feature requests too!
                    <br /> <a target="_blank" href="https://github.com/odpf/siren/issues"> Go to issues </a>

                  </div>)
              },
              {
                title: 'GitHub Discussions',
                content: (
                  <div>
                    For help and questions about best practices, join our GitHub discussions. Browse and ask questions.
                    <br /><a target="_blank" href="https://github.com/odpf/siren/discussions"> Go to discussions </a>

                  </div>)
              }
            ]}
          />
        </Container>
      </main>
    </Layout>
  )
}