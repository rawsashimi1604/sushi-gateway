import React, { useState } from 'react'
import Tag from '../typography/Tag'
import NormalText from '../typography/Text'
import PluginDropdown from './PluginDropdown'
import { IoIosArrowDown, IoIosArrowUp } from 'react-icons/io'

function ServiceDropdown() {

    const [isClicked, setIsClicked] = useState(false)

    return (
        <div>
            {/* Service metadata */}
            <div className="bg-neutral-200 px-4 py-2 rounded-lg shadow-sm w-[80%]">
                <div
                    className="flex items-center justify-between"
                    onClick={() => setIsClicked((prev) => !prev)}
                >
                    <div className="flex items-center gap-2">
                        <h1 className="text-md tracking-wide">service_name</h1>
                        <h1 className="text-xs italic text-neutral-800 mt-0.5">service_name</h1>
                    </div>

                    {isClicked ? <IoIosArrowUp /> : <IoIosArrowDown />}
                </div>
                {isClicked && (
                    <div className="flex flex-col items-start gap-2 text-sm">
                        <div className="flex items-center gap-2 text-sm">
                            <Tag value="name" />
                            <NormalText text="some_service_name" />
                        </div>
                        <div className="flex items-center gap-2 text-sm">
                            <Tag value="base_path" />
                            <NormalText text="/some_base_path" />
                        </div>


                        <div className="flex items-center gap-2 text-sm">
                            <Tag value="protocol" />
                            <NormalText text="http" />
                        </div>

                        <div className="flex flex-col justify-center items-start  text-sm">
                            <Tag value="upstreams" />
                            <div className='flex flex-col bg-neutral-100 px-4 py-2 mt-2'>
                                <NormalText text="localhost:8080" />
                                <NormalText text="localhost:8081" />
                            </div>

                        </div>

                        <div className='w-full'>
                            <div className="mb-2">
                                <Tag value="plugins" />
                            </div>

                            {/* Some dropdown for plugin design */}
                            <PluginDropdown name="http_log" data={{
                                "name": "http_log",
                                "enabled": true,
                                "data": {
                                    "http_endpoint": "http://localhost:8003/v1/log",
                                    "method": "POST",
                                    "content_type": "application/json"
                                }
                            }} />
                        </div>
                    </div>)
                }
            </div>


        </div>
    )
}

export default ServiceDropdown