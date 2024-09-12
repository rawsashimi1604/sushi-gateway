interface SubtitleProps {
  text: string;
}

function Subtitle({ text }: SubtitleProps) {
  return <div className="tracking-wider text-gray-500 text-sm">{text}</div>;
}

export default Subtitle;
