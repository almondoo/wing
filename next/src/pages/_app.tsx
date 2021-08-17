import '../styles/reset.css';
import { SWRConfig, Revalidator } from 'swr';
import Layout from '../components/identities/_app/layout/index';
import type { AppProps } from 'next/app';
import type { RevalidatorOptions } from 'swr/dist/types';

function MyApp({ Component, pageProps }: AppProps): JSX.Element {
  return (
    <Layout>
      <SWRConfig
        value={{
          dedupingInterval: 600000,
          revalidateOnFocus: false,
          onErrorRetry: (
            error: { status: number },
            key: string,
            config,
            revalidate: Revalidator,
            { retryCount }: Required<RevalidatorOptions>,
          ): void => {
            //- 404は再試行しない
            if (error.status === 404) return;

            if (retryCount >= 5) return;

            //- 5秒後に再試行
            setTimeout(() => revalidate({ retryCount }), 3000);
          },
        }}
      >
        <Component {...pageProps} />
      </SWRConfig>
    </Layout>
  );
}

export default MyApp;
