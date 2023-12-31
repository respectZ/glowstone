package component

// Code generated by scrapper/entityBp.ts;

// Defines this entity's ability to trade with players.  
type EconomyTradeTable struct {
  // Determines when the mob transforms, if the trades should be converted when the new mob has a economy_trade_table. When the trades are converted, the mob will generate a new trade list with their new trade table, but then it will try to convert any of the same trades over to have the same enchantments and user data. For example, if the original has a Emerald to Enchanted Iron Sword (Sharpness 1), and the new trade also has an Emerald for Enchanted Iron Sword, then the enchantment will be Sharpness 1.
  ConvertTradesEconomy bool `json:"convert_trades_economy,omitempty"`
  // How much should the discount be modified by when the player has cured the Zombie Villager. Can be specified as a pair of numbers (When use_legacy_price_formula is true this is the low-tier trade discount and high-tier trade discount, otherwise it is the minor positive gossip and major positive gossip.)
  CuredDiscount *float64 `json:"cured_discount,omitempty"`
  // Name to be displayed while trading with this entity
  DisplayName string `json:"display_name,omitempty"`
  // Used in legacy prices to determine how much should Demand be modified by when the player has the Hero of the Village mob effect
  HeroDemandDiscount *int `json:"hero_demand_discount,omitempty"`
  // The max the discount can be modified by when the player has cured the Zombie Villager. Can be specified as a pair of numbers (When use_legacy_price_formula is true this is the low-tier trade discount and high-tier trade discount, otherwise it is the minor positive gossip and major positive gossip.)
  MaxCuredDiscount *float64 `json:"max_cured_discount,omitempty"`
  // The max the discount can be modified by when the player has cured a nearby Zombie Villager. Only used when use_legacy_price_formula is true, otherwise max_cured_discount (low) is used.
  MaxNearbyCuredDiscount *int `json:"max_nearby_cured_discount,omitempty"`
  // How much should the discount be modified by when the player has cured a nearby Zombie Villager
  NearbyCuredDiscount *int `json:"nearby_cured_discount,omitempty"`
  // Used to determine if trading with entity opens the new trade screen
  NewScreen bool `json:"new_screen,omitempty"`
  // Determines if the trades should persist when the mob transforms. This makes it so that the next time the mob is transformed to something with a trade_table or economy_trade_table, then it keeps their trades.
  PersistTrades bool `json:"persist_trades,omitempty"`
  // Show an in game trade screen when interacting with the mob.
  ShowTradeScreen *bool `json:"show_trade_screen,omitempty"`
  // File path relative to the resource pack root for this entity's trades
  Table string `json:"table,omitempty"`
  // Determines whether the legacy formula is used to determines the trade prices.
  UseLegacyPriceFormula bool `json:"use_legacy_price_formula,omitempty"`
}
