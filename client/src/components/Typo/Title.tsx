import clsx from "clsx";

interface TitleProps {
  as?: "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  className?: string;
  children: React.ReactNode;
}

export const Title: React.FC<TitleProps> = ({
  as = "h3",
  className,
  children,
}) => {
  const Comp = as;
  return (
    <Comp
      className={clsx(
        "font-bold tracking-tight text-white",
        "title--" + as,
        className
      )}
    >
      {children}
    </Comp>
  );
};
