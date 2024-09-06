interface TagProps {
  value: string;
  bgColor?: string;
}

function Tag({ value, bgColor = "bg-gray-100" }: TagProps) {
  return (
    // TODO: change color scheme
    <span className={`px-2.5 py-1 rounded-xl ${bgColor}`}>{value}</span>
  );
}

export default Tag;
