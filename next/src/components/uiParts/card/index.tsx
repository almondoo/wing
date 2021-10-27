import { ReactNode } from 'react';
import Style from './style';

type Props = {
  children: ReactNode;
};

const Card = ({ children, ...props }: Props): JSX.Element => {
  return <Style.Card {...props}>{children}</Style.Card>;
};

export default Card;
