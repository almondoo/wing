import { useState, ReactNode, ChangeEvent } from 'react';
import Style from './style';

type Props = {
  type?: string;
  value: string | number;
  isFloat?: boolean;
  isRequired?: boolean;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  children: ReactNode;
};
const TextField = ({
  type = 'text',
  value,
  onChange,
  isFloat = false,
  isRequired = false,
  children,
  ...props
}: Props): JSX.Element => {
  const [inputType, setInputType] = useState(type);

  const handleChangeVisible = () => {
    if (inputType === 'password') {
      setInputType('text');
    } else {
      setInputType('password');
    }
  };

  return (
    <Style.Wrapper>
      <Style.Field>
        <Style.Input onChange={onChange} {...props} />
        <Style.Label isFloat={isFloat} isInput={value}>
          {children}
          {isRequired ? <Style.Required>*</Style.Required> : ''}
        </Style.Label>
        {type === 'password' ? <Style.EyeIcon onClick={handleChangeVisible} /> : ''}
      </Style.Field>
    </Style.Wrapper>
  );
};

export default TextField;
