import clsx from "clsx";

interface ContainerProps {
  children: React.ReactNode;
  className?: string;
}
export function Container(props: ContainerProps) {
  return (
    <div className={clsx("sm:px-8", props.className)}>
      <div className="mx-auto max-w-7xl lg:px-8">
        <div className={clsx("relative px-4 sm:px-8 lg:px-12")}>
          <div className="mx-auto max-w-2xl lg:max-w-5xl">{props.children}</div>
        </div>
      </div>
    </div>
  );
}
