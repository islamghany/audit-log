//import Link from 'next/link'
import clsx from "clsx";

enum Variant {
  PRIMARY,
  SECONDARY,
}
const styles: Record<Variant, string> = {
  [Variant.PRIMARY]:
    "rounded-full bg-primary py-2 px-4 text-sm font-semibold text-white hover:bg-indigo-600 active:bg-indigo-700 focus:outline-none focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-300/50",
  [Variant.SECONDARY]:
    "rounded-full bg-slate-800 py-2 px-4 text-sm font-medium text-white hover:bg-slate-700 active:text-slate-400 focus:outline-none focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white/50",
};

interface ButtonProps extends React.HTMLAttributes<HTMLButtonElement> {
  variant?: Variant;
  isLoading?: boolean;
}
export function Button({
  variant = Variant.PRIMARY,
  isLoading = false,
  className,
  ...props
}: ButtonProps) {
  return <button className={clsx(styles[variant], className)} {...props} />;
}
