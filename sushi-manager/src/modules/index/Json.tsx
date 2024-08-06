import Header from "../../components/typography/Header";
import ReactJson from "react-json-view";

interface JsonProps {
  data: any;
}

function Json({ data }: JsonProps) {
  return (
    <div>
      <Header text="Json" align="left" size="md" />
      <div className="bg-neutral-200 px-4 py-4 rounded-lg shadow-sm w-[80%]">
        <ReactJson src={data} />
      </div>
    </div>
  );
}

export default Json;
