import { ReactNode, ChangeEvent } from 'react';
import Style from './style';

type Props = {
  rows?: number;
  value: string;
  isRequired?: boolean;
  onChange?: (v: ChangeEvent<HTMLTextAreaElement>) => void;
  children: ReactNode;
};

// テキストエリアs
const TextArea = ({
  rows = 4,
  value,
  isRequired = false,
  onChange,
  children,
  ...props
}: Props): JSX.Element => {
  const height = `${rows * 26 + 20}px`;

  return (
    <Style.Wrapper>
      <Style.Field height={height}>
        <Style.Input value={value} height={height} onChange={onChange} {...props} />
        <Style.Label isInput={value}>
          {children}
          {isRequired ? <Style.Required>*</Style.Required> : ''}
        </Style.Label>
      </Style.Field>
    </Style.Wrapper>
  );
};

export default TextArea;
