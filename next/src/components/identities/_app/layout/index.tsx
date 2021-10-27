import React, { ReactNode } from 'react';
import Style from './style';
import Head from 'next/head';

const Layout = ({ children }: { children: ReactNode }): JSX.Element => {
  return (
    <div>
      <Head>
        <link rel="icon" href="/favicon.png" type="image/png" />
      </Head>
      <Style.Wrapper>{children}</Style.Wrapper>
    </div>
  );
};

export default Layout;
