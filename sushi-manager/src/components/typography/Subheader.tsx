type SubheaderProps = {
  text: string;
  align?: "left" | "center" | "right";
  size?: "xs" | "sm" | "md" | "lg";
};

function Subheader({ text, align = "center", size = "lg" }: SubheaderProps) {
  const sizeMap = {
    xs: "text-lg",
    sm: "text-xl ",
    md: "text-2xl ",
    lg: "text-3xl ",
  };

  const textAlignmentMap = {
    left: "text-left",
    center: "text-center",
    right: "text-right",
  };

  return (
    <h2
      className={`font-sans tracking-wider ${sizeMap[size]}
      ${textAlignmentMap[align]}
    `}
    >
      {text}
    </h2>
  );
}

export default Subheader;
