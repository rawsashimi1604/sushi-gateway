interface TagProps {
  value: string;
  className?: string;
  bgColor?: string;
}

function Tag({ value, className, bgColor = "bg-gray-100" }: TagProps) {
  return (
    // TODO: change color scheme
    <span className={`px-2.5 py-1 rounded-xl ${bgColor} ${className}`}>
      {value}
    </span>
  );
}

export default Tag;
