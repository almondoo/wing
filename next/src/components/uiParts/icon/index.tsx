import Style from './style';

type Props = {
  src: string;
  alt: string;
  size: number;
};

const Icon = ({ src, alt, size, ...props }: Props): JSX.Element => {
  return (
    <Style.Wrapper size={size} {...props}>
      <Style.Image src={src} alt={alt} layout="fill" objectFit="cover" />
    </Style.Wrapper>
  );
};

export default Icon;
