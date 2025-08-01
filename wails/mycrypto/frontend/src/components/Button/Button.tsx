import styles from './Button.module.css';

type ButtonVariant = 'primary' | 'secondary' | 'tertiary';

export type ButtonProps = {
  onClick: () => void;
  children: React.ReactNode;
  variant?: ButtonVariant;
};

export default function Button(props: ButtonProps): JSX.Element {
  let variantStyle = '';

  switch (props.variant) {
    case 'primary':
    default:
      break;

    case 'secondary':
      variantStyle = styles.secondary;
      break;

    case 'tertiary':
      variantStyle = styles.tertiary;
      break;
  }
  return (
    <button className={`${styles.button} ${variantStyle}`} onClick={() => props.onClick()}>
      {props.children}
    </button>
  );
}
