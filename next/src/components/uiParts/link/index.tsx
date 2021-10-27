import { ReactNode, MouseEventHandler, forwardRef } from 'react';
import Style from './style';
import NextLink from 'next/link';

type ForwardProps = {
  children?: ReactNode;
  href?: string;
  onClick?: MouseEventHandler;
  isUnderline: boolean;
};

const Item = forwardRef<HTMLAnchorElement, ForwardProps>(function Item(
  { children, href, onClick, isUnderline, ...props },
  ref,
) {
  return (
    <Style.Link href={href} onClick={onClick} ref={ref} isUnderline={isUnderline} {...props}>
      {children}
    </Style.Link>
  );
});

type Props = {
  children: React.ReactNode;
  href: string;
  isUnderline?: boolean;
};

const Link = ({ children, href, isUnderline = true, ...props }: Props): JSX.Element => {
  return (
    <NextLink href={href} {...props} passHref>
      <Item isUnderline={isUnderline}>{children}</Item>
    </NextLink>
  );
};

export default Link;
