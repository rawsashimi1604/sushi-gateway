import PluginDropdown from "../../components/sushi-gateway/PluginDropdown";
import Header from "../../components/typography/Header";
import Tag from "../../components/typography/Tag";
import NormalText from "../../components/typography/Text";

interface GlobalProps {
  data: any;
}

function Global({ data }: GlobalProps) {
  return (
    <section className="">
      <Header text="Global" align="left" size="md" />

      {/* Gateway metadata */}
      <div className="bg-neutral-200 px-4 py-4 rounded-lg shadow-sm w-[80%]">
        <div className="flex items-center gap-2 text-sm">
          <Tag value="name" />
          <NormalText text={data?.name} />
        </div>

        {/* Gateway plugins */}
        <div>
          <div className="mb-2">
            <Tag value="plugins" />
          </div>
          <div className="flex flex-col gap-2">
            {data?.plugins.map((plugin: any) => {
              return (
                <PluginDropdown
                  key={plugin?.name}
                  name={plugin?.name}
                  data={plugin}
                />
              );
            })}
          </div>
        </div>
      </div>
    </section>
  );
}

export default Global;
