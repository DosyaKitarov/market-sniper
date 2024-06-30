package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//type Product struct {
//	ASIN          string `json:"asin"`
//	Name          string `json:"name"`
//	Brand         string `json:"brand"`
//	Price         string `json:"price"`
//	PreviousPrice string `json:"previous_price"`
//	ChangeDate    string `json:"change_date"`
//}
//
//func (p *Product) FormatProduct(asin, inputJSON string)([]byte, error {
//	var data map[string]interface{}
//	err := json.Unmarshal([]byte(inputJSON), &data)
//	if err != nil {
//		return nil, err
//	}
//
//	p.ASIN = asin
//	name, _ := data["name"].(string)
//	p.Name = name
//	brand, _ := data["brand"].(string)
//	p.Brand = brand
//	price, _ := data["pricing"].(string)
//	p.Price = price
//
//}
//
//func main() {
//	asin := "B0BVDJBDXF"
//	p := Product{}
//	json := "{\"name\":\"Speedo Unisex-child Swim Snorkel Dive Mask Anti-Fog with Nose Cover Adventure Series\",\"product_information\":{\"brand\":\"Speedo\",\"color\":\"M Blue\",\"material\":\"Polycarbonate; 100% silicone\",\"closure_type\":\"Pull On\",\"age_range_description\":\"Big Kid\",\"included_components\":\"Dive Mask\",\"strap_type\":\"Adjustable\",\"item_weight\":\"0.35 Pounds\",\"global_trade_identification_number\":\"00827782641579\",\"upc\":\"827782641579\",\"item_package_dimensions_l_x_w_x_h\":\"9 x 7 x 3 inches\",\"package_weight\":\"0.16 Kilograms\",\"item_dimensions_lx_wx_h\":\"9 x 8 x 4 inches\",\"brand_name\":\"Speedo\",\"model_name\":\"Swim Snorkel Dive Mask Anti-Fog with Nose Cover Adventure Series\",\"suggested_users\":\"unisex-child\",\"number_of_items\":\"1\",\"manufacturer\":\"Warnaco Swimwear - Speedo Equipment\",\"part_number\":\"8-7530333420-1SZ\",\"style\":\"Swim Snorkel Dive Mask Anti-Fog with Nose Cover Adventure Series\",\"size\":\"One Size\",\"sport_type\":\"Swimming\",\"asin\":\"B002MUPY3E\",\"customer_reviews\":{\"ratings_count\":8218,\"stars\":4.6},\"best_sellers_rank\":[\"#855 in Sports & Outdoors (See Top 100 in Sports & Outdoors)\",\"#54 in Swimming Goggles\"],\"date_first_available\":\"August 25, 2009\"},\"brand\":\"Visit the Speedo Store\",\"brand_url\":\"https://www.amazon.com/stores/Speedo/page/012FFA9E-E5FB-44FA-8A23-EEA4DB36C659?ref_=ast_bln&store_ref=bl_ast_dp_brandLogo_sto\",\"full_description\":\"They'll be able to focus on creating underwater adventures with this full-coverage mask designed for smaller faces. Anti-fog lenses make it easy to keep a sharp lookout for treasures, while a silicone skirt provides a snug fit that keeps water out of their eyes for a more comfortable swim.\",\"pricing\":\"$18.31\",\"list_price\":\"$22.50\",\"shipping_price\":\"FREE\",\"availability_status\":\"In Stock\",\"is_coupon_exists\":false,\"images\":[\"https://m.media-amazon.com/images/I/31D7BjJkN0L.jpg\",\"https://m.media-amazon.com/images/I/41Lr-aIIcpL.jpg\",\"https://m.media-amazon.com/images/I/41b6Crw3AFL.jpg\",\"https://m.media-amazon.com/images/I/51+0njE8rjL.jpg\",\"https://m.media-amazon.com/images/I/41dWamwy0hL.jpg\",\"https://m.media-amazon.com/images/I/5137ww65mfL.jpg\",\"https://m.media-amazon.com/images/I/51wtC+6S29L.jpg\"],\"product_category\":\"Sports & Outdoors › Sports › Water Sports › Swimming › Goggles\",\"average_rating\":4.6,\"feature_bullets\":[\"They'll be able to focus on creating underwater adventures with this full-coverage mask designed for smaller faces. Anti-fog lenses make it easy to keep a sharp lookout for treasures, while a silicone skirt provides a snug fit that keeps water out of their eyes for a more comfortable swim.\",\"100% silicone skirt for ultimate comfort\",\"Speedo exclusive Anti Fog Max coating for a clear view\",\"Dual colored translucent mask with Speedo's trademarked Speed Fit headstrap system for quick adjustability\",\"Polycarbonate lens for great clarity underwater\",\"Perfect for snorkeling and recreational swimillimetering for kids or adults with smaller, narrower faces\"],\"total_reviews\":8218,\"customization_options\":{\"color\":[{\"is_selected\":true,\"value\":\"M Blue\",\"image\":\"https://m.media-amazon.com/images/I/31D7BjJkN0L.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B00HYH42U6/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"M Blue Sea\",\"image\":\"https://m.media-amazon.com/images/I/41GVBLi0JEL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B0BRCKYS8Q/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"M Bright Blue/Clear\",\"image\":\"https://m.media-amazon.com/images/I/415gGe25U3L.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B09R3XN5JJ/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"M Green Gecko\",\"image\":\"https://m.media-amazon.com/images/I/31aOlhtV7AL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B00536Y58E/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"M Pink Frost\",\"image\":\"https://m.media-amazon.com/images/I/41T2WoUgDuL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B09R3ZP3VF/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"M Spectra Yellow/Clear\",\"image\":\"https://m.media-amazon.com/images/I/31t8jRFplRL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B002MUTX5E/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"M Speedo Green\",\"image\":\"https://m.media-amazon.com/images/I/31e-zyvSWNL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B07494HZJQ/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"MS Blue Sea\",\"image\":\"https://m.media-amazon.com/images/I/41nY30kEvkL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B09R43XSPV/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"MS Green Gecko/Clear\",\"image\":\"https://m.media-amazon.com/images/I/31Jv+NOdXHL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B0748Z6FD4/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"MS Pink Frost\",\"image\":\"https://m.media-amazon.com/images/I/31uHSjsH2WL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B09R44FYM8/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"MS Spectra Yellow/Clear\",\"image\":\"https://m.media-amazon.com/images/I/31hwlePdftL.jpg\"},{\"is_selected\":false,\"url\":\"https://www.amazon.com/dp/B0BRCMJ1C2/ref=twister_B0CFY34NK2?_encoding=UTF8&psc=1\",\"value\":\"MS Speedo Black/Clear\",\"image\":\"https://m.media-amazon.com/images/I/315cvdJUR1L.jpg\"}]},\"ships_from\":\"Amazon.com\",\"sold_by\":\"Amazon.com\",\"aplus_present\":true}"
//	err := p.FormatProduct(asin, json)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Println(p)
//}
