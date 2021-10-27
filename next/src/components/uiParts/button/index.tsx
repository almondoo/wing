import { ReactNode } from 'react';
import Style from './style';
import variable from '../../utils/variable';

type Props = {
  variant?: string;
  color?: string;
  fullWidth?: boolean;
  children: ReactNode;
  onClick?: () => void;
};

// ボタン
const Button = ({
  variant = 'contained',
  color = variable.color.primary,
  fullWidth = false,
  children,
  onClick,
  ...props
}: Props): JSX.Element => {
  return (
    <Style.Button
      variant={variant}
      fullWidth={fullWidth}
      color={color}
      onClick={onClick}
      {...props}
    >
      {children}
    </Style.Button>
  );
};

export default Button;
