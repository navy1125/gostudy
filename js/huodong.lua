
local function getTable(tblname)
	local request = HttpRequest:new()
	request:setRequestType(CCHttpRequest.kHttpPost)
	--request:setUrl("http://mbgame.bwgame.com.cn/"..gamename .. "/update")
	request:setUrl("http://180.168.197.87:28080/js/" .. game_name .. "/" .. tblname)
	request:send(function(_,response)
		if response:getResponseCode() == 200 and string.len(response:getResponseData()) > 0 then
			str = FileCrypto:DecodeBuffer(response:getResponseData())
			if str == nil then
				str = response:getResponseData()
			end
			local func,_ = loadstring(str)
			if func ~= nil then
				func()
			end
		end
	end)

end
function getTableConfig()
	local request = HttpRequest:new()
	request:setRequestType(CCHttpRequest.kHttpPost)
	--request:setUrl("http://mbgame.bwgame.com.cn/"..gamename .. "/update")
	request:setUrl("http://180.168.197.87:28080/js/" .. game_name .. "/TableConfig.lua")
	request:send(function(_,response)
		if response:getResponseCode() == 200 and string.len(response:getResponseData()) > 0 then
			str = FileCrypto:DecodeBuffer(response:getResponseData())
			if str == nil then
				str = response:getResponseData()
			end
			local func,_ = loadstring(str)
			if func ~= nil then
				func()
				for _,item in ipairs(TableConfig) do 
					if type(item) == "table" then
						getTable(item.updatefile)
					end
				end
			end
		end
	end)

end

local function checkUpdate()
	if g_lastUpdateTimeHuodong == nil or g_lastUpdateTimeHuodong < os.time() then
		require "script/gxlua/message"
		message.popOkMessage(600,100,"神秘活动开启,商城打折!")
		g_lastUpdateTimeHuodong = os.time()
		getTableConfig()
	end
end
checkUpdate()