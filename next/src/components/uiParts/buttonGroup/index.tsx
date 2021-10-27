import { ReactNode } from 'react';
import Style from './style';

type Props = {
  justifyContent?:
    | 'center'
    | 'flex-start'
    | 'flex-end'
    | 'space-between'
    | 'space-around'
    | 'space-evenly';
  children: ReactNode;
};

const ButtonGroup = ({ justifyContent = 'flex-start', children, ...props }: Props): JSX.Element => {
  return (
    <Style.ButtonWrap justifyContent={justifyContent} {...props}>
      {children}
    </Style.ButtonWrap>
  );
};

export default ButtonGroup;
