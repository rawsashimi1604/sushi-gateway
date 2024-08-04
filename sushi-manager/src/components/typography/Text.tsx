
interface NormalTextProps {
    text: string;
}

function NormalText({ text }: NormalTextProps) {
    return (
        <div className="tracking-wider">{text}</div>
    )
}

export default NormalText