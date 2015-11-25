
package definitions

func init(){
  add(`AskForPassword`, &defAskForPassword{})
}

type defAskForPassword struct{}

func (*defAskForPassword) String() string {
	return `
<interface>
  <object class="GtkDialog" id="AskForPassword">
    <property name="window-position">1</property>
    <property name="title" translatable="yes">Enter your password</property>
    <child internal-child="vbox">
      <object class="GtkBox" id="Vbox">
        <property name="margin">10</property>
        <child>
          <object class="GtkGrid" id="grid">
            <property name="margin">6</property>
            <property name="margin-bottom">10</property>
            <property name="row-spacing">12</property>
            <property name="column-spacing">6</property>
            <child>
              <object class="GtkLabel" id="accountMessage" >
                <property name="label" translatable="yes">Account</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                </attributes>
              </object>
              <packing>
                <property name="left-attach">0</property>
                <property name="top-attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="accountName" >
                <property name="halign">GTK_ALIGN_START</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="top-attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="passMessage" >
                <property name="label" translatable="yes">Password</property>
              </object>
              <packing>
                <property name="left-attach">0</property>
                <property name="top-attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="password">
                <property name="has-focus">true</property>
                <property name="visibility">false</property>
                <signal name="activate" handler="on_save_signal" />
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="top-attach">1</property>
              </packing>
            </child>
          </object>
        </child>
      </object>
    </child>
    <child type="action">
      <object class="GtkButton" id="button_cancel">
        <property name="label">Cancel</property>
        <signal name="clicked" handler="on_cancel_signal" />
      </object>
    </child>
    <child type="action">
      <object class="GtkButton" id="button_ok">
        <property name="label" translatable="yes">Connect</property>
        <property name="can-default">true</property>
        <signal name="clicked" handler="on_save_signal" />
      </object>
    </child>
  </object>
</interface>

`
}
