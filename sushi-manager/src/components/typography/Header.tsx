type HeaderProps = {
  text: string;
  align?: "left" | "center" | "right";
  className?: string;
  hasTextDecoration?: boolean;
  size?: "xs" | "sm" | "md" | "lg" | "xl" | "xxl";
};

function Header({
  text,
  align = "center",
  className,
  hasTextDecoration = false,
  size = "lg",
}: HeaderProps) {
  const sizeMap = {
    xs: "text-lg",
    sm: "text-xl",
    md: "text-2xl ",
    lg: "text-3xl ",
    xl: "text-4xl",
    xxl: "text-5xl",
  };

  const textAlignmentMap = {
    left: "text-left",
    center: "text-center",
    right: "text-right",
  };

  return (
    <h1
      className={`${sizeMap[size]} mb-2 tracking-wider font-lora font-light ${
        hasTextDecoration &&
        "underline decoration-gray-200 underline-offset-[6px] "
      }  
      ${textAlignmentMap[align]} ${className}
    `}
    >
      {text}
    </h1>
  );
}

export default Header;
