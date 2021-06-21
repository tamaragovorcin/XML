function getAuthHeader() {
	if (localStorage.getItem("expireTime") <= new Date().getTime()) {
		localStorage.removeItem("keyToken");
		localStorage.removeItem("keyRole");
		localStorage.removeItem("expireTime");
	}

	return `Bearer ${localStorage.getItem("keyToken")}`;
}

export default getAuthHeader;
